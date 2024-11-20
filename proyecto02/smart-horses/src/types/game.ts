export type Position = {
  row: number;
  col: number;
};

export type Horse = {
  position: Position;
  hasMultiplier: boolean;
};

export type BoardValue = {
  points?: number;
  multiplier?: boolean;
  horse?: 'white' | 'black';
};

export type Difficulty = 'BEGINNER' | 'AMATEUR' | 'EXPERT';

export interface GameState {
  board: BoardValue[][];
  whiteHorse: Horse;
  blackHorse: Horse;
  whiteScore: number;
  blackScore: number;
  currentPlayer: 'white' | 'black';
}

export type Move = {
  from: Position;
  to: Position;
};
