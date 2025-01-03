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
  playerColor: 'white' | 'black' | null;
  onPieceSelection: (color: 'white' | 'black' | null) => void;
}

const GameLayout: React.FC<GameLayoutProps> = ({
  gameState,
  difficulty,
  isGameOver,
  selectedSquare,
  onSquareClick,
  onDifficultyChange,
  onReset,
  playerColor,
  onPieceSelection,
}) => {
  const currentYear = new Date().getFullYear();
  return (
	<div className="min-h-screen bg-gray-900 text-white p-8 rounded-md">
      {!playerColor ? (
        <div className="piece-selection text-center bg-gray-800 p-6 gap-10 rounded-md w-full max-w-xl xl:max-w-6xl mx-auto flex flex-col items-center justify-center min-h-screen">
		<h2 className="text-2xl font-bold mb-6 text-white">Bienvenido a</h2>
		<h1 className="text-5xl font-bold text-center mb-8">Smart Horses</h1>
		<h2 className="text-2xl mb-6 text-white">Seleccione su color</h2>
		<div className="flex items-center justify-center gap-10 px-6 py-6">
		  <img
			src="SVG/green-horse-white.svg"
			alt="Pieza Blanca"
			onClick={() => onPieceSelection('white')}
			className="cursor-pointer max-h-40 max-w-full h-auto w-1/3 drop-shadow-lg"
		  />
		  <img
			src="SVG/orange-horse-black.svg"
			alt="Pieza Negra"
			onClick={() => onPieceSelection('black')}
			className="cursor-pointer max-h-40 max-w-full h-auto w-1/3 drop-shadow-lg"
		  />
		</div>
	  </div>
      ) : gameState ? (
      <div className="min-h-screen bg-gray-900 text-white p-8 rounded-md">
		  <div className="w-full xl:max-w-6xl xl:w-full mx-auto">
			  {/* Header */}
			  <h1 className="text-5xl font-bold text-center mb-8">Smart Horses</h1>
			  <div className="max-w-5xl mx-auto">
				  <div className="mb-8 text-center w-full gap-2">
					  <DifficultySelector
						  difficulty={difficulty}
						  onSelect={onDifficultyChange}
					  />
					  <div className="piece-selection text-center bg-gray-800 p-2 gap-1 xl:gap-20 rounded-md w-full max-w-xl xl:max-w-6xl mx-auto flex flex-row items-center justify-center">
						<h2 className="text-l xl:text-xl font-bold text-white">Seleccione su color</h2>
						<div className="flex items-center justify-center gap-1 xl:gap-20 px-1 py-1">
						<img
							src="SVG/green-horse-white.svg"
							alt="Pieza Blanca"
							onClick={() => onPieceSelection('white')}
							className={`cursor-pointer max-h-24 max-w-full h-auto w-1/3 ${playerColor === 'white' ? 'glow-white' : ''}`}
						/>
						<img
							src="SVG/orange-horse-black.svg"
							alt="Pieza Negra"
							onClick={() => onPieceSelection('black')}
							className={`cursor-pointer max-h-24 max-w-full h-auto w-1/3 ${playerColor === 'black' ? 'glow-white' : ''}`}
						/>
						</div>
	  				  </div>
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
					  <div className="flex flex-col min-[520px]:flex-row xl:flex-col w-full xl:space-y6 space-x-0 min-[520px]:space-x-4 xl:space-x-0 gap-y-4">
						  <div className="w-full min-[520px]:w-1/2 flex flex-col gap-4 xl:w-full">
						  		{/* Game status */}	
								{isGameOver && (
									<GameStatus
										isGameOver={isGameOver}
										winner={gameState.whiteScore > gameState.blackScore ? 'white' : 
																		gameState.blackScore > gameState.whiteScore ? 'black' : 'draw'}
									/>	
								)}

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
									  Turno: {gameState.currentPlayer === playerColor ? 'Jugador' : 'IA'}
								  </p>
							  </div>
						  </div>
						  
						  {/* Instrucciones */}
						  <div className="bg-gray-800 text-left rounded-lg p-4 w-full min-[520px]:w-1/2 xl:w-full xl:mt-12">
							  <h3 className="font-bold mb-3">Instrucciones:</h3>
							  <ul className="space-y-2 text-gray-300">
								  <li>• Haz click en el caballo para seleccionarlo</li>
								  <li>• Haz click en una casilla válida para mover</li>
								  <li>• Obtén puntos al llegar a casillas numeradas</li>
								  <li>• x2 multiplica los siguientes puntos que obtengas</li>
							  </ul>
						  </div>

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
		  <p className="read-the-docs mt-8">
                © {currentYear}. Molina, JS; Narvaéz, JC; Pacheco, CD; Puyo, JE. Inteligencia Artificial; Ingeniería en sistemas; EISC<br/> Todos los derechos reservados.
            </p>
      </div> ) : (
        <div>Loading...</div>
      )}
    </div>
  );
};

export default GameLayout;
