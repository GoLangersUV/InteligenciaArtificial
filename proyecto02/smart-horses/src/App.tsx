import { useState, useEffect } from 'react';
import { Difficulty } from './types/game';
import { useGame } from './hooks/useGame';
import GameLayout from './components/Layout/GameLayout';
import { SoundManager } from './utils/soundManager';
import './App.css';

function App() {
  const [difficulty, setDifficulty] = useState<Difficulty>('BEGINNER');
  const [playerColor, setPlayerColor] = useState<'white' | 'black' | null>(null);

  const {
    gameState,
    isGameOver,
    selectedSquare,
    handleSquareClick,
    resetGame,
  } = useGame(difficulty, playerColor);

  const handleDifficultyChange = (newDifficulty: Difficulty) => {
    setDifficulty(newDifficulty);
  };

  const handlePieceSelection = (color: 'white' | 'black' | null) => {
    setPlayerColor(color);
  };

  useEffect(() => {
    SoundManager.loadSounds();
  }, []);

  return (
    <GameLayout
      gameState={gameState}
      difficulty={difficulty}
      isGameOver={isGameOver}
      selectedSquare={selectedSquare}
      onSquareClick={handleSquareClick}
      onDifficultyChange={handleDifficultyChange}
      onReset={resetGame}
      playerColor={playerColor}
      onPieceSelection={handlePieceSelection}
  />
  );
}

export default App;
