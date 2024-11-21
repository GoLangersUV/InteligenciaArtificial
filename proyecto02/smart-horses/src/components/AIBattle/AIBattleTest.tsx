import React, { useState } from 'react';
import { AIBattleSystem, BattleProgress } from '../../logic/ai/battleSystem';

const AIBattleTest: React.FC = () => {
  const [isRunning, setIsRunning] = useState(false);
  const [results, setResults] = useState<Record<string, {ai1Wins: number, ai2Wins: number, draws: number}> | null>(null);
  const [progress, setProgress] = useState<BattleProgress | null>(null);

  const difficulties = ['BEGINNER', 'AMATEUR', 'EXPERT'];
  const difficultyNames = {
    'BEGINNER': 'Principiante',
    'AMATEUR': 'Amateur',
    'EXPERT': 'Experto'
  };

  const runBattles = async () => {
    setIsRunning(true);
    setResults(null);
    setProgress(null);
    
    try {
      const battleResults = await AIBattleSystem.runAllBattles((progress) => {
        setProgress(progress);
      });
      setResults(battleResults);
    } finally {
      setIsRunning(false);
    }
  };

  return (
    <div className="p-4 bg-gray-800 rounded-lg max-w-4xl mx-auto">
      <div className="flex items-center justify-between mb-6">
        <div className="text-left">
          <h2 className="text-xl font-bold text-white">Pruebas AI vs AI</h2>
          <p className="mt-2 text-sm text-gray-300">
            Resultados de enfrentamientos entre IAs.
          </p>
        </div>
        <button
          onClick={runBattles}
          disabled={isRunning}
          className="px-4 py-2 bg-blue-600 text-white rounded disabled:bg-gray-600 hover:bg-blue-700 transition-colors font-semibold"
        >
          {isRunning ? 'Ejecutando...' : 'Iniciar pruebas'}
        </button>
      </div>

      {isRunning && progress && (
        <div className="mb-6 space-y-4">
          <div className="space-y-2">
            <div className="flex justify-between text-sm text-gray-300">
              <span>Progreso: {Math.round((progress.completedMatches / progress.totalMatches) * 100)}%</span>
              <span>{progress.completedMatches}/{progress.totalMatches}</span>
            </div>
            <div className="w-full bg-gray-700 rounded-full h-2">
              <div
                className="bg-blue-600 h-2 rounded-full transition-all duration-300"
                style={{ width: `${(progress.completedMatches / progress.totalMatches) * 100}%` }}
              />
            </div>
          </div>

          <div className="bg-gray-700 p-3 rounded-lg">
            <p className="text-sm font-medium mb-1 text-gray-200">{progress.currentMatchup}</p>
            <div className="flex justify-between text-sm text-gray-300">
              <span>IA1: {progress.currentAI1Score}</span>
              <span>IA2: {progress.currentAI2Score}</span>
            </div>
          </div>
        </div>
      )}

      {results && (
        <div className="mt-6">
          <div className="overflow-x-auto">
            <table className="w-full text-left border-separate border-spacing-x-3 border-spacing-y-2">
              <thead>
                <tr>
                  <th className="text-sm font-medium text-gray-400 w-24">IA1 vs IA2</th>
                  {difficulties.map(diff => (
                    <th key={diff} className="text-sm font-medium text-gray-200 w-24">
                      {difficultyNames[diff as keyof typeof difficultyNames]}
                    </th>
                  ))}
                </tr>
              </thead>
              <tbody>
                {difficulties.map((ai1Diff) => (
                  <tr key={ai1Diff}>
                    <td className="text-sm font-medium text-gray-200">
                      {difficultyNames[ai1Diff as keyof typeof difficultyNames]}
                    </td>
                    {difficulties.map((ai2Diff) => {
                      const key = `${ai1Diff}_vs_${ai2Diff}`;
                      const result = results[key];
                      return (
                        <td key={ai2Diff} className="text-sm">[&nbsp;
                          <span className="text-green-400">{result.ai1Wins}</span>
                          {' '}
                          <span className="text-red-400">{result.ai2Wins}</span>
                          {' '}
                          <span className="text-gray-400">{result.draws}</span>
                        &nbsp;]
                        </td>
                      );
                    })}
                  </tr>
                ))}
              </tbody>
            </table>
          </div>
          <div className="mt-4 text-xs text-center text-gray-400">
            <span className="text-green-400">Victorias IA1</span>
            {' · '}
            <span className="text-red-400">Victorias IA2</span>
            {' · '}
            <span className="text-gray-400">Empates</span>
          </div>
        </div>
      )}
    </div>
  );
};

export default AIBattleTest;
