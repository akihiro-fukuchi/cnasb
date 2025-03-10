from podcastfy.client import generate_podcast
import os

# OpenAI API key (replace with your actual key)
# os.environ["OPENAI_API_KEY"] = "sk-xxx"

# Configuration for Japanese podcast generation
custom_config = {
    "word_count": 20000,
    "conversation_style": ["casual"],
    "podcast_name": "AI技術ポッドキャスト",
    "creativity": 1,
    "output_language": "ja",
    "default_tts_model": "openai",
    "model_name": "openai/o3-mini",
    "text_to_speech": {
        "output_directories": {
            "audio": ".",
            "transcripts": "."
        }
    }
}

def main():
    try:
        # Read the transcript file
        with open("cline/transcript-formatted.txt", "r", encoding="utf-8") as f:
            content = f.read()

        print("Generating podcast from transcript...")
        
        # Generate podcast
        generate_podcast(
            text=content,
            conversation_config=custom_config,
            tts_model="openai",
            llm_model_name="openai/o1",
            api_key_label="OPENAI_API_KEY",
            longform=True
        )
        
        print("Podcast generation completed. Check the current directory for the audio file.")
        
    except FileNotFoundError:
        print("Error: transcript file not found")
    except Exception as e:
        print(f"An error occurred: {str(e)}")

if __name__ == "__main__":
    main()
