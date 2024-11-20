import React from 'react';
import Board from '../Board/Board';
import DifficultySelector from '../GameInfo/DifficultySelector';
import ScoreBoard from '../GameInfo/ScoreBoard';
import GameStatus from '../GameInfo/GameStatus';
import { GameState, Position, Difficulty } from '../../types/game';

interface GameLayoutProps {
  gameState: GameState;
  difficulty: Difficulty;
  isGameOver: boolean;
  selectedSquare: Position | null;
  onSquareClick: (position: Position) => void;
  onDifficultyChange: (difficulty: Difficulty) => void;
  onReset: () => void;
}

const GameLayout: React.FC<GameLayoutProps> = ({
  gameState,
  difficulty,
  isGameOver,
  selectedSquare,
  onSquareClick,
  onDifficultyChange,
  onReset
}) => {
  // Determinar el ganador
  const getWinner = () => {
    if (!isGameOver) return undefined;
    if (gameState.whiteScore > gameState.blackScore) return 'white';
    if (gameState.blackScore > gameState.whiteScore) return 'black';
    return 'draw';
  };

  return (
    <div className="container mx-auto p-4">
      {/* Header */}
      <div className="mb-6 text-center">
        <h1 className="text-3xl font-bold mb-4">Smart Horses</h1>
        <DifficultySelector
          difficulty={difficulty}
          onSelect={onDifficultyChange}
        />
      </div>

      {/* Main game area */}
      <div className="flex flex-col lg:flex-row gap-6">
        {/* Game board */}
        <div className="flex-grow">
          <Board
            gameState={gameState}
            onSquareClick={onSquareClick}
            selectedSquare={selectedSquare}
          />
        </div>

        {/* Game info sidebar */}
        <div className="w-full lg:w-64 space-y-4">
          <ScoreBoard
            whiteScore={gameState.whiteScore}
            blackScore={gameState.blackScore}
          />
          
          <GameStatus
            isGameOver={isGameOver}
            winner={getWinner()}
          />

          <button
            onClick={onReset}
            className="w-full px-4 py-2 bg-blue-500 text-white rounded hover:bg-blue-600 transition-colors"
          >
            Reiniciar Juego
          </button>

          {/* Turno actual */}
          <div className="p-4 bg-white shadow rounded">
            <p className="text-center font-medium">
              Turno: {gameState.currentPlayer === 'white' ? 'IA' : 'Jugador'}
            </p>
          </div>

          {/* Instrucciones básicas */}
          <div className="p-4 bg-white shadow rounded">
            <h3 className="font-bold mb-2">Instrucciones:</h3>
            <ul className="text-sm space-y-1">
              <li>• Haz click en tu caballo para seleccionarlo</li>
              <li>• Haz click en una casilla válida para mover</li>
              <li>• Obtén puntos al llegar a casillas numeradas</li>
              <li>• x2 multiplica los siguientes puntos que obtengas</li>
            </ul>
          </div>
        </div>
      </div>
    </div>
  );
};

export default GameLayout;
