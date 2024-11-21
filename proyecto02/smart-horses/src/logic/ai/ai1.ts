import { GameStateManager } from '../gameState';
import { GameUtilities } from '../utilities';
import { Position } from '../../types/game';

export function evaluatePositionAI1(state: GameStateManager): number {
  // IA1: Estrategia agresiva centrada en capturar puntos y multiplicadores
  
  const baseScore = state.blackScore - state.whiteScore;
  let positionScore = 0;
  
  // Cache para distancias
  const distanceCache = new Map<string, number>();
  const getDistanceFromCache = (from: Position, to: Position): number => {
    const key = `${from.row},${from.col}-${to.row},${to.col}`;
    if (!distanceCache.has(key)) {
      distanceCache.set(key, GameUtilities.getManhattanDistance(from, to));
    }
    return distanceCache.get(key)!;
  };

  // Evaluar proximidad a puntos altos y multiplicadores
  let bestPointNearby = 0;
  let isNearMultiplier = false;
  
  for (let row = 0; row < 8; row++) {
    for (let col = 0; col < 8; col++) {
      if (state.board[row][col].points) {
        const points = state.board[row][col].points!;
        const distance = getDistanceFromCache(
          state.blackHorse.position,
          { row, col }
        );
        // Priorizar puntos altos
        if (points > bestPointNearby && distance < 4) {
          bestPointNearby = points;
        }
        // Bonus por estar cerca de puntos
        positionScore += points / (distance + 1);
      }
      
      if (state.board[row][col].multiplier) {
        const distance = getDistanceFromCache(
          state.blackHorse.position,
          { row, col }
        );
        if (distance < 3) isNearMultiplier = true;
      }
    }
  }

  // Bonus por estar cerca de puntos altos
  if (bestPointNearby > 5) {
    positionScore += bestPointNearby * 2;
  }
  
  // Bonus por estar cerca de multiplicador cuando hay puntos altos disponibles
  if (isNearMultiplier && bestPointNearby > 5) {
    positionScore += 30;
  }
  
  return (baseScore * 10) + positionScore;
}
