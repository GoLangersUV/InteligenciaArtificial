import React, { useState } from 'react';
import { AIBattleSystem, BattleProgress } from '../../logic/ai/battleSystem';

const AIBattleTest: React.FC = () => {
  const [isRunning, setIsRunning] = useState(false);
  const [results, setResults] = useState<string>('');
  const [progress, setProgress] = useState<BattleProgress | null>(null);

  const runBattles = async () => {
    setIsRunning(true);
    setResults('');
    setProgress(null);
    
    try {
      const battleResults = await AIBattleSystem.runAllBattles((progress) => {
        setProgress(progress);
      });
      
      const table = AIBattleSystem.generateResultsTable(battleResults);
      setResults(table);
    } finally {
      setIsRunning(false);
    }
  };

  return (
    <div className="p-4 bg-gray-800 rounded-lg">
      <div className="flex justify-between items-center mb-6">
        <h2 className="text-xl font-bold">Pruebas AI vs AI</h2>
        <button
          onClick={runBattles}
          disabled={isRunning}
          className="px-4 py-2 bg-blue-500 text-white rounded disabled:bg-gray-400 hover:bg-blue-600 transition-colors"
        >
          {isRunning ? 'Ejecutando pruebas...' : 'Iniciar pruebas'}
        </button>
      </div>

      {isRunning && progress && (
        <div className="mb-6 space-y-4">
          <div className="space-y-2">
            <div className="flex justify-between text-sm text-gray-300">
              <span>Progreso Total: {Math.round((progress.completedMatches / progress.totalMatches) * 100)}%</span>
              <span>{progress.completedMatches} de {progress.totalMatches} partidas</span>
            </div>
            <div className="w-full bg-gray-700 rounded-full h-2.5">
              <div
                className="bg-blue-500 h-2.5 rounded-full transition-all duration-300"
                style={{ 
                  width: `${(progress.completedMatches / progress.totalMatches) * 100}%` 
                }}
              />
            </div>
          </div>

          <div className="bg-gray-700 p-4 rounded-lg">
            <h3 className="text-sm font-medium mb-2">Partida Actual:</h3>
            <p className="text-lg font-bold mb-2">{progress.currentMatchup}</p>
            <div className="flex justify-between text-sm">
              <span>AI1: {progress.currentAI1Score}</span>
              <span>AI2: {progress.currentAI2Score}</span>
            </div>
          </div>
        </div>
      )}

      {results && (
        <div className="mt-6">
          <h3 className="font-bold mb-2">Resultados Finales:</h3>
          <div className="bg-gray-900 p-4 rounded overflow-x-auto">
            <pre className="text-sm text-gray-300">
              {results}
            </pre>
          </div>
        </div>
      )}
    </div>
  );
};

export default AIBattleTest;
