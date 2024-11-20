import { useState, useEffect, useCallback, useRef } from 'react';
import { GameStateManager } from '../logic/gameState';
import { Minimax } from '../logic/minimax';
import { Position, Difficulty } from '../types/game';
import { DIFFICULTIES } from '../constants/gameConstants';
import { evaluatePositionAI1 } from '../logic/ai/ai1';

export function useGame(difficulty: Difficulty) {
  const [gameState, setGameState] = useState<GameStateManager>(new GameStateManager());
  const [selectedSquare, setSelectedSquare] = useState<Position | null>(null);
  const [isGameOver, setIsGameOver] = useState(false);
  const [aiThinking, setAiThinking] = useState(false);
  const isInitialMove = useRef(true);
  const minimax = new Minimax(evaluatePositionAI1, DIFFICULTIES[difficulty].depth);

  const checkGameOver = useCallback((state: GameStateManager) => {
    if (!state.hasPointsRemaining()) {
      console.log('Game Over - No points remaining');
      setIsGameOver(true);
    }
  }, []);

  const makeAIMove = useCallback(async () => {
    console.log('makeAIMove called with state:', gameState);

    if (gameState.currentPlayer !== 'white' || aiThinking) {
      console.log('makeAIMove cancelled', { currentPlayer: gameState.currentPlayer, aiThinking });
      return;
    }
    
    setAiThinking(true);
    
    try {
      const currentState = gameState.clone();
      const move = minimax.getBestMove(currentState);
      
      if (move) {
        const success = currentState.makeMove(move.from, move.to);
        
        if (success) {
          console.log('AI move successful, updating state');
          setGameState(currentState);
          checkGameOver(currentState);
        }
      }
    } catch (error) {
      console.error('Error in AI move:', error);
    } finally {
      setAiThinking(false);
    }
}, [gameState, minimax, checkGameOver, aiThinking]);

  const handleSquareClick = useCallback((position: Position) => {
    console.log('handleSquareClick', {
      position,
      currentPlayer: gameState.currentPlayer,
      aiThinking,
      isGameOver,
      selectedSquare
    });

    if (aiThinking || isGameOver || gameState.currentPlayer !== 'black') {
      console.log('Click ignored - invalid state');
      return;
    }

    if (!selectedSquare) {
      if (
        position.row === gameState.blackHorse.position.row && 
        position.col === gameState.blackHorse.position.col
      ) {
        console.log('Black horse selected');
        setSelectedSquare(position);
      } else {
        console.log('Invalid selection - not black horse');
      }
    } else {
      console.log('Attempting move from', selectedSquare, 'to', position);
      const newGameState = gameState.clone();
      const moveSuccess = newGameState.makeMove(selectedSquare, position);
      
      if (moveSuccess) {
        console.log('Player move successful');
        setGameState(newGameState); // Actualizar el estado
        setSelectedSquare(null);
        checkGameOver(newGameState);
        
        if (!isGameOver) {
          console.log('Scheduling AI move with updated state');
          setTimeout(() => {
            setGameState(prevState => {
              console.log('Current state before AI move:', prevState);
              const aiState = prevState.clone();
              const aiMove = minimax.getBestMove(aiState);
              
              if (aiMove) {
                console.log('AI attempting move:', aiMove);
                const aiMoveSuccess = aiState.makeMove(aiMove.from, aiMove.to);
                if (aiMoveSuccess) {
                  console.log('AI move successful');
                  return aiState;
                }
              }
              return prevState;
            });
          }, 500);
        }
      } else {
        console.log('Invalid move attempted');
        setSelectedSquare(null);
      }
    }
}, [gameState, selectedSquare, isGameOver, aiThinking, checkGameOver, minimax]);

  useEffect(() => {
    console.log('Game initialization started');
    const initGame = () => {
      const newGameState = new GameStateManager();
      setGameState(newGameState);
      setSelectedSquare(null);
      setIsGameOver(false);
      setAiThinking(false);
      isInitialMove.current = true;
    };

    initGame();
    console.log('Game initialized');
  }, [difficulty]);

  useEffect(() => {
    if (isInitialMove.current && gameState.currentPlayer === 'white' && !aiThinking) {
      console.log('Initiating first AI move');
      isInitialMove.current = false;
      new Promise(resolve => setTimeout(resolve, 500))
        .then(() => {
          console.log('Executing first AI move');
          return makeAIMove();
        })
        .catch(error => {
          console.error('Error in first AI move:', error);
        });
    }
  }, [gameState, makeAIMove, aiThinking]);

  const resetGame = useCallback(() => {
    console.log('Game reset requested');
    const newGameState = new GameStateManager();
    setGameState(newGameState);
    setSelectedSquare(null);
    setIsGameOver(false);
    setAiThinking(false);
    isInitialMove.current = true;
    console.log('Game reset completed');
  }, []);

  return {
    gameState,
    isGameOver,
    selectedSquare,
    handleSquareClick,
    resetGame,
  };
}
