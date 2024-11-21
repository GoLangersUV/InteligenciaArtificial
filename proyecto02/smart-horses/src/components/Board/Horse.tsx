import React from 'react';

interface HorseProps {
  type: 'white' | 'black';
}

const Horse: React.FC<HorseProps> = ({ type }) => {
  const symbol = type === 'white' ? '♞': '♞';
  const color = type === 'white' ? 'text-red' : 'text-black';

  return (
    <div className={`text-3xl drop-shadow-[0_1.2px_1.2px_rgba(0,0,0,0.8)] ${color}`}>{symbol}</div>
  );
};

export default Horse;
