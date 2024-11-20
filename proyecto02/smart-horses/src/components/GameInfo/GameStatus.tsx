interface GameStatusProps {
  isGameOver: boolean;
  winner: 'white' | 'black' | 'draw';
}

const GameStatus: React.FC<GameStatusProps> = ({ isGameOver, winner }) => {
  if (!isGameOver) return null;

  const message = winner === 'draw' ? '¡Empate!' : `Ganador: ${winner === 'white' ? 'Blancas' : 'Negras'}`;

  return (
    <div className="text-xl font-bold text-center p-4 bg-blue-100 rounded">
      {message}
    </div>
  );
};

export default GameStatus;
