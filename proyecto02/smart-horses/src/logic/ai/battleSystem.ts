import { GameStateManager } from '../gameState';
import { Difficulty } from '../../types/game';
import { DIFFICULTIES } from '../../constants/gameConstants';
import { Minimax } from '../minimax';
import { evaluatePositionAI1 } from './ai1';
import { evaluatePositionAI2 } from './ai2';

export type BattleProgress = {
  totalMatches: number;
  completedMatches: number;
  currentMatchup: string;
  currentAI1Score: number;
  currentAI2Score: number;
  phase: 'RUNNING' | 'COMPLETED';
};

export type BattleResult = {
  ai1Wins: number;
  ai2Wins: number;
  draws: number;
};

type MatchResult = {
  winner: 'AI1' | 'AI2' | 'DRAW';
  finalScore: {
    ai1: number;
    ai2: number;
  };
};

export class AIBattleSystem {
  private static readonly MATCHES_PER_COMBINATION = 10;
  private static readonly TOTAL_MATCHES = 90; // 9 combinaciones x 10 partidas

  static async runAllBattles(
    onProgress: (progress: BattleProgress) => void
  ): Promise<Record<string, BattleResult>> {
    const difficulties: Difficulty[] = ['BEGINNER', 'AMATEUR', 'EXPERT'];
    const results: Record<string, BattleResult> = {};
    let completedMatches = 0;

    for (const ai1Difficulty of difficulties) {
      for (const ai2Difficulty of difficulties) {
        const currentMatchup = `${ai1Difficulty} vs ${ai2Difficulty}`;
        const key = `${ai1Difficulty}_vs_${ai2Difficulty}`;
        
        onProgress({
          totalMatches: this.TOTAL_MATCHES,
          completedMatches,
          currentMatchup,
          currentAI1Score: 0,
          currentAI2Score: 0,
          phase: 'RUNNING'
        });

        const result = await this.runBattleSeries(
          ai1Difficulty, 
          ai2Difficulty,
          (score) => {
            onProgress({
              totalMatches: this.TOTAL_MATCHES,
              completedMatches: completedMatches + 1,
              currentMatchup,
              currentAI1Score: score.ai1,
              currentAI2Score: score.ai2,
              phase: 'RUNNING'
            });
          }
        );

        results[key] = result;
        completedMatches += this.MATCHES_PER_COMBINATION;
        
        console.log(`Resultados ${currentMatchup}:`);
        console.log(`AI1 victorias: ${result.ai1Wins}`);
        console.log(`AI2 victorias: ${result.ai2Wins}`);
        console.log(`Empates: ${result.draws}`);
        console.log('------------------------');
      }
    }

    onProgress({
      totalMatches: this.TOTAL_MATCHES,
      completedMatches,
      currentMatchup: 'Completado',
      currentAI1Score: 0,
      currentAI2Score: 0,
      phase: 'COMPLETED'
    });

    return results;
  }

  private static async runBattleSeries(
    ai1Difficulty: Difficulty,
    ai2Difficulty: Difficulty,
    onMatchProgress: (score: { ai1: number; ai2: number }) => void
  ): Promise<BattleResult> {
    const result: BattleResult = {
      ai1Wins: 0,
      ai2Wins: 0,
      draws: 0
    };

    for (let i = 0; i < this.MATCHES_PER_COMBINATION; i++) {
      const matchResult = await this.runSingleMatch(
        ai1Difficulty, 
        ai2Difficulty,
        onMatchProgress
      );
      
      if (matchResult.winner === 'AI1') result.ai1Wins++;
      else if (matchResult.winner === 'AI2') result.ai2Wins++;
      else result.draws++;
      
      await new Promise(resolve => setTimeout(resolve, 0));
    }

    return result;
  }

  private static async runSingleMatch(
  ai1Difficulty: Difficulty,
  ai2Difficulty: Difficulty,
  onProgress: (score: { ai1: number; ai2: number }) => void
): Promise<MatchResult> {
  const gameState = new GameStateManager('white');
  const ai1 = new Minimax(evaluatePositionAI1, DIFFICULTIES[ai1Difficulty].depth);
  const ai2 = new Minimax(evaluatePositionAI2, DIFFICULTIES[ai2Difficulty].depth);

  let moveCount = 0;
  const MAX_MOVES = 100; // Límite de seguridad

  console.log('Iniciando partida:');
  console.log('Estado inicial:', {
    whitePos: gameState.whiteHorse.position,
    blackPos: gameState.blackHorse.position,
    currentPlayer: gameState.currentPlayer
  });

  while (gameState.hasPointsRemaining() && moveCount < MAX_MOVES) {
    moveCount++;
    
    // Verificar estado antes del movimiento
    const preScore = {
      white: gameState.whiteScore,
      black: gameState.blackScore
    };

    // Seleccionar IA según el turno (AI2 blancas, AI1 negras)
    const isWhiteTurn = gameState.currentPlayer === 'white';
    const currentAI = isWhiteTurn ? ai2 : ai1;
    
    console.log(`\nTurno ${moveCount}:`, 
      isWhiteTurn ? 'AI2 (blancas)' : 'AI1 (negras)',
      `Puntuación actual - Blancas: ${preScore.white}, Negras: ${preScore.black}`
    );

    // Obtener y validar movimiento
    const move = currentAI.getBestMove(gameState);
    if (!move) {
      console.error('No se encontró movimiento válido');
      break;
    }

    console.log('Movimiento elegido:', move);

    // Realizar movimiento
    const moveSuccess = gameState.makeMove(move.from, move.to);
    if (!moveSuccess) {
      console.error('Movimiento inválido');
      break;
    }

    // Verificar si hubo cambio en la puntuación
    const postScore = {
      white: gameState.whiteScore,
      black: gameState.blackScore
    };

    if (postScore.white !== preScore.white || postScore.black !== preScore.black) {
      console.log('¡Puntos capturados!', {
        whiteDelta: postScore.white - preScore.white,
        blackDelta: postScore.black - preScore.black
      });
    }

    // Reportar progreso
    onProgress({
      ai1: gameState.blackScore,
      ai2: gameState.whiteScore
    });

    await new Promise(resolve => setTimeout(resolve, 0));
  }

  console.log('\nFin de partida:', {
    movimientos: moveCount,
    puntuaciónBlancas: gameState.whiteScore,
    puntuaciónNegras: gameState.blackScore,
  });

  return {
    winner: gameState.blackScore > gameState.whiteScore ? 'AI1' :
            gameState.whiteScore > gameState.blackScore ? 'AI2' : 'DRAW',
    finalScore: {
      ai1: gameState.blackScore,
      ai2: gameState.whiteScore
    }
  };
}

  static generateResultsTable(results: Record<string, BattleResult>): string {
    let table = '| IA1 vs IA2 | Principiante | Amateur | Experto |\n';
    table += '|------------|--------------|---------|----------|\n';

    const difficulties: Difficulty[] = ['BEGINNER', 'AMATEUR', 'EXPERT'];
    
    for (const ai1Diff of difficulties) {
      let row = `| ${ai1Diff} |`;
      
      for (const ai2Diff of difficulties) {
        const key = `${ai1Diff}_vs_${ai2Diff}`;
        const result = results[key];
        row += ` [${result.ai1Wins} ${result.ai2Wins} ${result.draws}] |`;
      }
      
      table += row + '\n';
    }

    return table;
  }
}
