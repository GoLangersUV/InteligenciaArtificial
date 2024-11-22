import { useState, useEffect, useCallback, useRef } from 'react';
import { GameStateManager } from '../logic/gameState';
import { Minimax } from '../logic/minimax';
import { Position, Difficulty } from '../types/game';
import { DIFFICULTIES } from '../constants/gameConstants';
import { evaluatePositionAI1 } from '../logic/ai/ai1';
import { SoundManager } from '../utils/soundManager';

export function useGame(difficulty: Difficulty, playerColor: 'white' | 'black' | null) {
  const [gameState, setGameState] = useState<GameStateManager>(new GameStateManager('white'));
  const [selectedSquare, setSelectedSquare] = useState<Position | null>(null);
  const [isGameOver, setIsGameOver] = useState(false);
  const [aiThinking, setAiThinking] = useState(false);
  const minimaxRef = useRef(new Minimax(evaluatePositionAI1, DIFFICULTIES[difficulty].depth));
  const isInitialMove = useRef(true);

//  const checkGameOver = useCallback((state: GameStateManager) => {
//    if (!state.hasPointsRemaining()) {
//      console.log('Game Over - No points remaining');
//      setIsGameOver(true);
//    }
//  }, []);

const makeAIMove = useCallback(() => {
  setGameState((currentState) => {
    if (currentState.currentPlayer === playerColor || aiThinking) {
      return currentState;
    }

    setAiThinking(true);

    try {
      const stateCopy = currentState.clone();
      const move = minimaxRef.current.getBestMove(stateCopy);

      if (move) {
        const success = stateCopy.makeMove(move.from, move.to);

        if (success) {
          // Handle move success...
          if (!stateCopy.hasPointsRemaining()) {
            setIsGameOver(true);
          }
          return stateCopy;
        }
      }

      return currentState;
    } catch (error) {
      console.error('Error in AI move:', error);
      return currentState;
    } finally {
      setAiThinking(false);
    }
  });
}, [aiThinking, playerColor]);


  const handleSquareClick = useCallback((position: Position) => {
    if (aiThinking || isGameOver || gameState.currentPlayer !== playerColor) {
      console.log('Click ignored - invalid state');
      return;
    }
  
    if (!selectedSquare) {
      if ( playerColor === 'white' && 
        position.row === gameState.whiteHorse.position.row && 
        position.col === gameState.whiteHorse.position.col
      ) {
        setSelectedSquare(position);
      }
      if ( playerColor === 'black' && 
        position.row === gameState.blackHorse.position.row && 
        position.col === gameState.blackHorse.position.col
      ) {
        setSelectedSquare(position);
      }

    } else {
      const newGameState = gameState.clone();
      const moveSuccess = newGameState.makeMove(selectedSquare, position);
      
      if (moveSuccess) {
        const toSquareValue = newGameState.board[position.row][position.col];
        const hasReward = toSquareValue.points || toSquareValue.multiplier;
  
        if (hasReward) {
          SoundManager.playSound('onReward');
        } else {
          SoundManager.playSound('onBoard');
        }
  
        setGameState(newGameState);
        setSelectedSquare(null);
        
        if (!newGameState.hasPointsRemaining()) {
          setIsGameOver(true);
        } else {
          // Programar el movimiento de la IA después de un breve delay
          setTimeout(makeAIMove, 500);
        }
      } else {
        setSelectedSquare(null);
      }
    }
  }, [gameState, selectedSquare, isGameOver, aiThinking, makeAIMove, playerColor]);
  

  // Efecto para inicializar el juego
  useEffect(() => {
    if (playerColor) {
      const newGameState = new GameStateManager('white'); // White always starts
      setGameState(newGameState);
      // Reset other states...
      isInitialMove.current = true;
  
      minimaxRef.current = new Minimax(evaluatePositionAI1, DIFFICULTIES[difficulty].depth);
  
      if (playerColor === 'black') {
        // AI makes the first move
        setTimeout(makeAIMove, 500);
      }
    }
  }, [difficulty, playerColor, makeAIMove]);

  // Efecto para el primer movimiento de la IA si le toca
  useEffect(() => {
    if (
      isInitialMove.current &&
      gameState.currentPlayer !== playerColor &&
      !aiThinking
    ) {
      isInitialMove.current = false;
      setTimeout(makeAIMove, 500);
    }
  }, [gameState, makeAIMove, aiThinking, playerColor]);

  const resetGame = useCallback(() => {
    if (playerColor) {
      const newGameState = new GameStateManager('white'); // White always starts
      setGameState(newGameState);
      setSelectedSquare(null);
      setIsGameOver(false);
      setAiThinking(false);
      isInitialMove.current = true;
  
      if (playerColor === 'black') {
        // AI makes the first move
        setTimeout(makeAIMove, 500);
      }
    }
  }, [playerColor, makeAIMove]);

  return {
    gameState,
    isGameOver,
    selectedSquare,
    handleSquareClick,
    resetGame,
  };
}
