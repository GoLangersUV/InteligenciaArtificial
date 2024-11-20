import React from 'react';
import { BoardValue, Position } from '../../types/game';
import Horse from './Horse';

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
        aspect-square w-full
        flex items-center justify-center
        relative
        ${isLightSquare ? 'bg-blue-200' : 'bg-gray-400'}
        ${isSelected ? 'ring-2 ring-blue-500' : ''}
        hover:bg-blue-100 cursor-pointer
        transition-colors
      `}
      onClick={onClick}
    >
      <div className="absolute inset-0 flex items-center justify-center">
        {value.points && (
          <span className="font-bold text-xl text-black z-10">
            {value.points}
          </span>
        )}
        {value.multiplier && (
          <span className="font-bold text-lg text-green-600 z-10">
            x2
          </span>
        )}
        {value.horse && (
          <div className="z-20">
            <Horse type={value.horse} />
          </div>
        )}
      </div>
    </div>
  );
};

export default Square;
