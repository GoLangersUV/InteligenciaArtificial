import React from 'react';

interface ScoreBoardProps {
  whiteScore: number;
  blackScore: number;
}

const ScoreBoard: React.FC<ScoreBoardProps> = ({ whiteScore, blackScore }) => {
  return (
    <div className="p-4 bg-white shadow rounded">
      <h2 className="text-xl font-bold mb-4">Puntajes</h2>
      <div className="space-y-2">
        <div className="flex justify-between">
          <span>Caballo Blanco:</span>
          <span className="font-bold">{whiteScore}</span>
        </div>
        <div className="flex justify-between">
          <span>Caballo Negro:</span>
          <span className="font-bold">{blackScore}</span>
        </div>
      </div>
    </div>
  );
};

export default ScoreBoard;
