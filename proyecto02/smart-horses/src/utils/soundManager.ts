// src/utils/soundManager.ts
export class SoundManager {
    private static sounds: { [key: string]: HTMLAudioElement } = {};
  
    static loadSounds() {
      this.sounds['onBoard'] = new Audio('/audio/on-board.mp3');
      this.sounds['onReward'] = new Audio('/audio/on-reward.mp3');
  
      this.sounds['onBoard'].volume = 0.5;
      this.sounds['onReward'].volume = 0.7;
    }
  
    static playSound(soundName: 'onBoard' | 'onReward') {
      const sound = this.sounds[soundName];
      if (sound) {
        sound.currentTime = 0;
        sound.play();
      }
    }
  }
  