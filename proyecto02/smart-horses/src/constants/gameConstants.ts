import { Difficulty } from '../types/game';

export const BOARD_SIZE = 8;
export const TOTAL_POINTS_SQUARES = 10;
export const TOTAL_MULTIPLIER_SQUARES = 4;

export const DIFFICULTIES: Record<Difficulty, { name: string; depth: number }> = {
  BEGINNER: { name: 'Principiante', depth: 2 },
  AMATEUR: { name: 'Amateur', depth: 4 },
  EXPERT: { name: 'Experto', depth: 6 }
} as const;

export const POINTS_RANGE = {
  MIN: 1,
  MAX: 10
} as const;

export const KNIGHT_MOVES = [
  [-2, -1], [-2, 1],
  [-1, -2], [-1, 2],
  [1, -2], [1, 2],
  [2, -1], [2, 1]
] as const;
