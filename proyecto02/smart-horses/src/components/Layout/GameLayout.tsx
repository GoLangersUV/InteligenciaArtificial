// src/components/Layout/GameLayout.tsx
import React from 'react';
import Board from '../Board/Board';
import DifficultySelector from '../GameInfo/DifficultySelector';
import ScoreBoard from '../GameInfo/ScoreBoard';
import GameStatus from '../GameInfo/GameStatus';
import AIBattleTest from '../AIBattle/AIBattleTest';
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
  return (
    <div className="min-h-screen bg-gray-900 text-white p-8">
      <div className="max-w-7xl mx-auto">
        {/* Header */}
        <h1 className="text-5xl font-bold text-center mb-8">Smart Horses</h1>
        
        <div className="mb-8 text-center">
          <DifficultySelector
            difficulty={difficulty}
            onSelect={onDifficultyChange}
          />
        </div>

        <div className="flex justify-center gap-8">
          {/* Game board */}
          <div className="flex-shrink-0">
            <Board
              gameState={gameState}
              onSquareClick={onSquareClick}
              selectedSquare={selectedSquare}
            />
          </div>

          {/* Game info sidebar */}
          <div className="w-64 space-y-6">
            <ScoreBoard
              whiteScore={gameState.whiteScore}
              blackScore={gameState.blackScore}
            />
            
            <button
              onClick={onReset}
              className="w-full py-3 px-4 bg-blue-600 text-white rounded-lg hover:bg-blue-700 transition-colors font-medium"
            >
              Reiniciar Juego
            </button>

            {/* Turno actual */}
            <p className="text-center font-medium">
              Turno: {gameState.currentPlayer === 'black' ? 'IA' : 'Jugador'}
            </p>
           

            {/* Instrucciones */}
            <div className="bg-gray-800 text-left rounded-lg p-4">
              <h3 className="font-bold mb-3">Instrucciones:</h3>
              <ul className="space-y-2 text-gray-300">
                <li>• Haz click en tu caballo para seleccionarlo</li>
                <li>• Haz click en una casilla válida para mover</li>
                <li>• Obtén puntos al llegar a casillas numeradas</li>
                <li>• x2 multiplica los siguientes puntos que obtengas</li>
              </ul>
            </div>

            {isGameOver && (
              <GameStatus
                isGameOver={isGameOver}
                winner={gameState.whiteScore > gameState.blackScore ? 'white' : 
                       gameState.blackScore > gameState.whiteScore ? 'black' : 'draw'}
              />
            )}
          </div>
        </div>

        {/* Sección de pruebas AI vs AI */}
        <div className="mt-12 border-t-2 border-gray-800 pt-8">
          <div className="max-w-3xl mx-auto">
            <h2 className="text-2xl font-bold text-center mb-6">Pruebas de IA vs IA</h2>
            <AIBattleTest />
          </div>
        </div>
      </div>
    </div>
  );
};

export default GameLayout;
