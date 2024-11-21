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
      <div className="min-h-screen bg-gray-900 text-white p-8 rounded-md">
		  <div className=" w-[720px] xl:max-w-6xl xl:w-full mx-auto">
			  {/* Header */}
			  <h1 className="text-5xl font-bold text-center mb-8">Smart Horses</h1>
			  <div className="max-w-5xl mx-auto">
				  <div className="mb-8 text-center">
					  <DifficultySelector
						  difficulty={difficulty}
						  onSelect={onDifficultyChange}
					  />
				  </div>
				  <div className="flex flex-col xl:flex-row justify-center gap-8">
					  {/* Game board */}
					  <div className="flex-shrink-0">
						  <Board
							  gameState={gameState}
							  onSquareClick={onSquareClick}
							  selectedSquare={selectedSquare}
						  />
					  </div>

					  {/* Game info sidebar */}
					  <div className="flex flex-row xl:flex-col w-full xl:space-y6 space-x-4 xl:space-x-0">
						  <div className="w-1/2 flex flex-col gap-4 xl:w-full">
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
							  <div className="bg-gray-800 rounded-lg p-4">
								  <p className="text-center font-medium">
									  Turno: {gameState.currentPlayer === 'white' ? 'Jugador' : 'IA'}
								  </p>
							  </div>
						  </div>
						  

            {/* Turno actual */}
            <p className="text-center font-medium">
              Turno: {gameState.currentPlayer === 'black' ? 'IA' : 'Jugador'}
            </p>
           
						  {/* Instrucciones */}
						  <div className="bg-gray-800 text-left rounded-lg p-4 w-1/2 xl:w-full xl:mt-12">
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
