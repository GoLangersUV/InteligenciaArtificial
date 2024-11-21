import { useState } from 'react';
import { Difficulty } from './types/game';
import { useGame } from './hooks/useGame';
import GameLayout from './components/Layout/GameLayout';
import './App.css';

function App() {
  const [difficulty, setDifficulty] = useState<Difficulty>('BEGINNER');
  const {
    gameState,
    isGameOver,
    selectedSquare,
    handleSquareClick,
    resetGame
  } = useGame(difficulty);

  const handleDifficultyChange = (newDifficulty: Difficulty) => {
    setDifficulty(newDifficulty);
    resetGame();
  };

  return (
		  <GameLayout
			  gameState={gameState}
			  difficulty={difficulty}
			  isGameOver={isGameOver}
			  selectedSquare={selectedSquare}
			  onSquareClick={handleSquareClick}
			  onDifficultyChange={handleDifficultyChange}
			  onReset={resetGame}
		  />
  );
}

export default App;
