import React from 'react';


interface SquareProps {
  position: Position;
  value: BoardValue;
  onClick: () => void;
  isSelected: boolean;
}

const Square: React.FC<SquareProps> = ({ position, value, onClick, isSelected }) => {
  const isLightSquare = (position.row + position.col) % 2 === 0;

  return (
    <div
      className={`
        w-16 h-16
        flex items-center justify-center
        relative
        ${isLightSquare ? 'bg-white' : 'bg-gray-200'}
        ${isSelected ? 'ring-2 ring-blue-500' : ''}|
      }`}
      onClick={onClick}
    >
      {value.horse && <Horse type={value.horse} />}
      {value.points && (
        <div className="absolute text-sm font-bold">{value.points}</div>
      )}
      {value.multiplier && (
        <div className="absolute text-green-500 font-bold">x2</div>
      )}
    </div>
  );
};

export default Square;
