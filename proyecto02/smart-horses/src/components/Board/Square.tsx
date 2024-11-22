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

  // Rutas de las imágenes según los valores
  const pointImageSrc = value.points ? `/SVG/${value.points}.svg` : null;
  const multiplierImageSrc = value.multiplier ? '/SVG/x2.svg' : null;

  return (
    <div
      className={`
        aspect-square w-full
        relative
        ${isLightSquare ? 'bg-orange-100' : 'bg-orange-300'}
        ${isSelected ? 'ring-2 ring-green-500 z-10' : ''}
        hover:bg-green-100 cursor-pointer
        transition-colors
      `}
      onClick={onClick}
    >
      {/* Contenedor de puntos y multiplicador */}
      <div className="absolute inset-0 flex items-center justify-center">
        {pointImageSrc && (
          <img
            src={pointImageSrc}
            alt={`Points ${value.points}`}
            className="w-full h-full object-contain z-10"
          />
        )}
        {multiplierImageSrc && (
          <img
            src={multiplierImageSrc}
            alt="x2 multiplier"
            className="w-full h-full object-contain z-10"
          />
        )}
      </div>

      {/* Contenedor del caballo */}
      {value.horse && (
        <div className="absolute inset-0 z-20">
          <Horse type={value.horse} />
        </div>
      )}
    </div>
  );
};

export default Square;
