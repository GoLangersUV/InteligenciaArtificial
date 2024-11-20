import React from 'react';
import { Difficulty } from '../../types/game';
import { DIFFICULTIES } from '../../constants/gameConstants';

interface DifficultySelectorProps {
  difficulty: Difficulty;
  onSelect: (difficulty: Difficulty) => void;
}

const DifficultySelector: React.FC<DifficultySelectorProps> = ({
  difficulty,
  onSelect,
}) => {
  return (
    <div className="mb-4">
      <label className="block text-sm font-medium text-white mb-2">
        Nivel de Dificultad
      </label>
      <select
        value={difficulty}
        onChange={(e) => onSelect(e.target.value as Difficulty)}
        className="mt-1 block w-full py-2 px-3 border border-gray-400 
                   bg-gray-800 rounded-lg shadow-sm 
                   focus:outline-none focus:ring-blue-200 focus:border-blue-200
                   text-white font-bold"
      >
        {Object.entries(DIFFICULTIES).map(([key, value]) => (
          <option key={key} value={key}>
            {value.name}
          </option>
        ))}
      </select>
    </div>
  );
};

export default DifficultySelector;
