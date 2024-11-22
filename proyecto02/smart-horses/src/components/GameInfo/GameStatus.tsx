interface GameStatusProps {
  isGameOver: boolean;
  winner: 'white' | 'black' | 'draw';
}

const GameStatus: React.FC<GameStatusProps> = ({ isGameOver, winner }) => {
  if (!isGameOver) return null;

  const message = winner === 'draw' ? '¡Empate!' : `Ganador: ${winner === 'white' ? 'Blanco' : 'Negro'}`;

  return (
    <div className="text-xl font-bold text-center p-4 mt-10 bg-yellow-400 rounded">
      {message}
    </div>
  );
};

export default GameStatus;
