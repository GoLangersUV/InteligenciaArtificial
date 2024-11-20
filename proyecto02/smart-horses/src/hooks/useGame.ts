import { useState, useEffect, useCallback } from 'react';
import { GameStateManager } from '../logic/gameState';
import { Minimax } from '../logic/minimax';
import { Position, Difficulty, GameState } from '../types/game';
import { DIFFICULTIES } from '../constants/gameConstants';
import { evaluatePositionAI1 } from '../logic/ai/ai1';
import { evaluatePositionAI2 } from '../logic/ai/ai2';

interface UseGameReturn {
  gameState: GameState;
  isGameOver: boolean;
  selectedSquare: Position | null;
  handleSquareClick: (position: Position) => void;
  resetGame: () => void;
}

export function useGame(difficulty: Difficulty): UseGameReturn {
  const [gameState, setGameState] = useState<GameStateManager>(new GameStateManager());
  const [selectedSquare, setSelectedSquare] = useState<Position | null>(null);
  const [isGameOver, setIsGameOver] = useState(false);
  const [aiThinking, setAiThinking] = useState(false);

  // Inicializar el minimax con la IA seleccionada
  const minimax = new Minimax(evaluatePositionAI1, DIFFICULTIES[difficulty].depth);

  // Resetear el juego
  const resetGame = useCallback(() => {
    const newGameState = new GameStateManager();
    setGameState(newGameState);
    setSelectedSquare(null);
    setIsGameOver(false);
  }, []);

  // Verificar si el juego ha terminado
  const checkGameOver = useCallback((state: GameStateManager) => {
    // El juego termina cuando no quedan casillas con puntos
    const hasPointsRemaining = state.board.some(row => 
      row.some(cell => cell.points !== undefined)
    );

    if (!hasPointsRemaining) {
      setIsGameOver(true);
    }
  }, []);

  // Manejar el click en una casilla
  const handleSquareClick = useCallback((position: Position) => {
    if (aiThinking || isGameOver || gameState.currentPlayer !== 'black') return;

    if (!selectedSquare) {
      // Si no hay casilla seleccionada, seleccionar esta si tiene el caballo del jugador
      if (
        position.row === gameState.blackHorse.position.row && 
        position.col === gameState.blackHorse.position.col
      ) {
        setSelectedSquare(position);
      }
    } else {
      // Si hay casilla seleccionada, intentar mover
      const newGameState = gameState.clone();
      if (newGameState.makeMove(selectedSquare, position)) {
        setGameState(newGameState);
        checkGameOver(newGameState);
        setSelectedSquare(null);
        
        // Turno de la IA
        if (!isGameOver) {
          makeAIMove(newGameState);
        }
      } else {
        setSelectedSquare(null);
      }
    }
  }, [gameState, selectedSquare, isGameOver, aiThinking]);

  // Realizar movimiento de la IA
  const makeAIMove = useCallback(async (currentState: GameStateManager) => {
    setAiThinking(true);
    
    // Simular un pequeño delay para que el movimiento de la IA no sea instantáneo
    await new Promise(resolve => setTimeout(resolve, 500));
    
    const move = minimax.getBestMove(currentState);
    if (move) {
      const newGameState = currentState.clone();
      newGameState.makeMove(move.from, move.to);
      setGameState(newGameState);
      checkGameOver(newGameState);
    }
    
    setAiThinking(false);
  }, [minimax, checkGameOver]);

  // Efecto para iniciar el juego
  useEffect(() => {
    resetGame();
  }, [difficulty]);

  return {
    gameState,
    isGameOver,
    selectedSquare,
    handleSquareClick,
    resetGame,
  };
}
