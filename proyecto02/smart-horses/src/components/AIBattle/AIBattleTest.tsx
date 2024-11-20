import React, { useState } from 'react';
import { AIBattleSystem } from '../../logic/ai/battleSystem';

const AIBattleTest: React.FC = () => {
  const [isRunning, setIsRunning] = useState(false);
  const [results, setResults] = useState<string>('');

  const runBattles = async () => {
    setIsRunning(true);
    try {
      const battleResults = await AIBattleSystem.runAllBattles();
      const table = AIBattleSystem.generateResultsTable(battleResults);
      setResults(table);
    } finally {
      setIsRunning(false);
    }
  };

  return (
    <div className="p-4">
      <h2 className="text-xl font-bold mb-4">Pruebas AI vs AI</h2>
      <button
        onClick={runBattles}
        disabled={isRunning}
        className="px-4 py-2 bg-blue-500 text-white rounded disabled:bg-gray-400"
      >
        {isRunning ? 'Ejecutando pruebas...' : 'Iniciar pruebas'}
      </button>

      {results && (
        <div className="mt-4">
          <h3 className="font-bold mb-2">Resultados:</h3>
          <pre className="bg-gray-100 p-4 rounded overflow-x-auto">
            {results}
          </pre>
        </div>
      )}
    </div>
  );
};

export default AIBattleTest;
