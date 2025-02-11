let words = [];
const wordDisplay = document.getElementById("word");
const inputField = document.getElementById("input");
const scoreDisplay = document.getElementById("score");
const timerDisplay = document.getElementById("timer");
const startButton = document.getElementById("startButton");

startButton.disabled = true;
fetch('https://raw.githubusercontent.com/kotofurumiya/pokemon_data/refs/heads/master/data/pokemon_data.json')
  .then(response => response.json())
  .then(data => {
    words = data;
    startButton.disabled = false;
  })
  .catch(error => {
    console.error("Error loading pokemon data:", error);
    startButton.disabled = false;
  });

let score = 0;
let currentWord = "";
let timeLeft = 60;
let intervalId;
let isGameRunning = false;

function katakanaToHiragana(str) {
  return str.replace(/[\u30a1-\u30f6]/g, m => String.fromCharCode(m.charCodeAt(0) - 0x60));
}

function pickWord() {
  let obj = words[Math.floor(Math.random() * 151)];
  currentWord = obj.name;
  wordDisplay.textContent = currentWord + " (No." + obj.no + ")";
}

function updateScore() {
  scoreDisplay.textContent = `スコア: ${score}`;
}

function updateTimer() {
  timerDisplay.textContent = `タイム: ${timeLeft}秒`;
}

function startGame() {
  if (!words || words.length === 0) return;
  score = 0;
  timeLeft = 60;
  inputField.value = "";
  inputField.disabled = false;
  pickWord();
  updateScore();
  updateTimer();
  isGameRunning = true;
  inputField.focus();
  intervalId = window.setInterval(() => {
    timeLeft--;
    updateTimer();
    if (timeLeft <= 0) {
      clearInterval(intervalId);
      gameOver();
    }
  }, 1000);
}

function gameOver() {
  clearInterval(intervalId);
  isGameRunning = false;
  inputField.disabled = true;
  const gameContainer = document.getElementById("game");
  gameContainer.style.display = "none";
  const overlay = document.getElementById("overlay");
  const finalScoreElement = document.getElementById("final-score");
  finalScoreElement.textContent = `最終スコア: ${score}`;
  overlay.classList.remove("hidden");
}

window.addEventListener("load", () => {
  const retryButton = document.getElementById("retry-button");
  const topButton = document.getElementById("top-button");
  if (retryButton) {
    retryButton.addEventListener("click", () => {
      const overlay = document.getElementById("overlay");
      overlay.classList.add("hidden");
      const gameContainer = document.getElementById("game");
      gameContainer.style.display = "block";
      startGame();
    });
  }
  if (topButton) {
    topButton.addEventListener("click", () => {
      const overlay = document.getElementById("overlay");
      overlay.classList.add("hidden");
      const gameContainer = document.getElementById("game");
      gameContainer.style.display = "block";
      score = 0;
      timeLeft = 60;
      updateScore();
      updateTimer();
      inputField.value = "";
      inputField.disabled = false;
      isGameRunning = false;
    });
  }
  document.addEventListener("keydown", (event) => {
    const overlay = document.getElementById("overlay");
    if ((event.key === "Enter" || event.key === " ") && !isGameRunning && overlay.classList.contains("hidden")) {
      startGame();
    } else if (event.key === "Escape" && isGameRunning) {
      clearInterval(intervalId);
      gameOver();
    }
  });
});

inputField.addEventListener("input", () => {
  if (katakanaToHiragana(inputField.value) === katakanaToHiragana(currentWord)) {
    score++;
    updateScore();
    inputField.value = "";
    pickWord();
  }
});

startButton.addEventListener("click", () => {
  if (!isGameRunning) {
    startGame();
  }
});
