import React from 'react';

interface HorseProps {
  type: 'white' | 'black';
}

const Horse: React.FC<HorseProps> = ({ type }) => {
  
  const horseImageSrc = type === 'white' ? '/SVG/green-horse.svg' : '/SVG/orange-horse.svg';

  return (
    <img
      src={horseImageSrc}
      alt={`${type} horse`}
      className="w-full h-full object-contain"
    />
  );
};

export default Horse;
