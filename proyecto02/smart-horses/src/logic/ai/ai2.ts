import { GameStateManager } from '../gameState';
import { GameUtilities } from '../utilities';

export function evaluatePositionAI2(state: GameStateManager): number {
  // IA2: Estrategia defensiva centrada en control de territorio y bloqueo
  
  const baseScore = state.whiteScore - state.blackScore;
  let positionScore = 0;

  // Evaluar movilidad
  const whiteMoves = state.getPossibleMoves(state.whiteHorse.position);
  const blackMoves = state.getPossibleMoves(state.blackHorse.position);
  
  // Bonus por tener más movimientos disponibles
  positionScore += (whiteMoves.length - blackMoves.length) * 5;

  // Evaluar control del centro
  const centerDistance = GameUtilities.getChebyshevDistance(
    state.whiteHorse.position,
    { row: 3.5, col: 3.5 }
  );
  positionScore += (7 - centerDistance) * 3;

  // Evaluar proximidad a puntos y oponente
  for (let row = 0; row < 8; row++) {
    for (let col = 0; col < 8; col++) {
      if (state.board[row][col].points) {
        const points = state.board[row][col].points!;
        const myDistance = GameUtilities.getManhattanDistance(
          state.whiteHorse.position,
          { row, col }
        );
        const opponentDistance = GameUtilities.getManhattanDistance(
          state.blackHorse.position,
          { row, col }
        );
        
        // Bonus por estar más cerca que el oponente
        if (myDistance < opponentDistance) {
          positionScore += points;
        }
        // Bonus por bloquear al oponente de puntos altos
        if (points > 5 && myDistance <= opponentDistance) {
          positionScore += points * 2;
        }
      }
    }
  }
  
  // Penalización por dejar al oponente cerca de multiplicador
  if (!state.blackHorse.hasMultiplier) {
    for (let row = 0; row < 8; row++) {
      for (let col = 0; col < 8; col++) {
        if (state.board[row][col].multiplier) {
          const opponentDistance = GameUtilities.getManhattanDistance(
            state.blackHorse.position,
            { row, col }
          );
          if (opponentDistance <= 2) {
            positionScore -= 20;
          }
        }
      }
    }
  }

  return (baseScore * 8) + positionScore;
}
