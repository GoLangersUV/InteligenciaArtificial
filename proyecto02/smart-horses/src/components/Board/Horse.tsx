import React from 'react';

interface HorseProps {
  type: 'white' | 'black';
}

const Horse: React.FC<HorseProps> = ({ type }) => {
  const symbol = type === 'white' ? '♘' : '♞';
  const color = type === 'white' ? 'text-white' : 'text-black';

  return (
    <div className={`text-3xl ${color}`}>{symbol}</div>
  );
};

export default Horse;
