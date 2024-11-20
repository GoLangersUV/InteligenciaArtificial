import { GameStateManager } from '../gameState';
import { Difficulty } from '../../types/game';
import { DIFFICULTIES } from '../../constants/gameConstants';
import { Minimax } from '../minimax';
import { evaluatePositionAI1 } from './ai1';
import { evaluatePositionAI2 } from './ai2';

type BattleResult = {
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
  private static readonly MATCHES_PER_COMBINATION = 1;

  static async runAllBattles(): Promise<Record<string, BattleResult>> {
    const difficulties: Difficulty[] = ['BEGINNER', 'AMATEUR', 'EXPERT'];
    const results: Record<string, BattleResult> = {};

    for (const ai1Difficulty of difficulties) {
      for (const ai2Difficulty of difficulties) {
        const key = `${ai1Difficulty}_vs_${ai2Difficulty}`;
        const result = await this.runBattleSeries(ai1Difficulty, ai2Difficulty);
        results[key] = result;
        
        console.log(`Resultados ${ai1Difficulty} vs ${ai2Difficulty}:`);
        console.log(`AI1 victorias: ${result.ai1Wins}`);
        console.log(`AI2 victorias: ${result.ai2Wins}`);
        console.log(`Empates: ${result.draws}`);
        console.log('------------------------');
      }
    }

    return results;
  }

  private static async runBattleSeries(
    ai1Difficulty: Difficulty,
    ai2Difficulty: Difficulty
  ): Promise<BattleResult> {
    const result: BattleResult = {
      ai1Wins: 0,
      ai2Wins: 0,
      draws: 0
    };

    for (let i = 0; i < this.MATCHES_PER_COMBINATION; i++) {
      const matchResult = await this.runSingleMatch(ai1Difficulty, ai2Difficulty);
      
      if (matchResult.winner === 'AI1') result.ai1Wins++;
      else if (matchResult.winner === 'AI2') result.ai2Wins++;
      else result.draws++;
      
      // Pequeña pausa para no bloquear el navegador
      await new Promise(resolve => setTimeout(resolve, 0));
    }

    return result;
  }

  private static async runSingleMatch(
    ai1Difficulty: Difficulty,
    ai2Difficulty: Difficulty
  ): Promise<MatchResult> {
    const gameState = new GameStateManager();
    const ai1 = new Minimax(evaluatePositionAI1, DIFFICULTIES[ai1Difficulty].depth);
    const ai2 = new Minimax(evaluatePositionAI2, DIFFICULTIES[ai2Difficulty].depth);

    while (gameState.hasPointsRemaining()) {
      const currentAI = gameState.currentPlayer === 'white' ? ai1 : ai2;
      const move = currentAI.getBestMove(gameState);
      
      if (!move) break;
      
      gameState.makeMove(move.from, move.to);
      await new Promise(resolve => setTimeout(resolve, 0));
    }

    return {
      winner: gameState.whiteScore > gameState.blackScore ? 'AI1' :
              gameState.blackScore > gameState.whiteScore ? 'AI2' : 'DRAW',
      finalScore: {
        ai1: gameState.whiteScore,
        ai2: gameState.blackScore
      }
    };
  }

  // Método para generar la tabla de resultados en formato markdown
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
