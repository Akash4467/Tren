Tren

A Local Voice-Controlled Desktop AI Assistant (Python + Speech Recognition + LLM)

Tren is a lightweight, local desktop assistant built in Python that listens to your voice, processes commands using an LLM, and speaks responses back. It aims to function as a fast, offline-friendly personal assistant for daily desktop tasks.

Features

Hands-free voice interaction
Speak naturally and Tren will capture, process, and understand your speech.

LLM-powered reasoning
Uses a language model (local or API-based) to understand and respond to user queries.

Text-to-Speech (TTS)
Responds with generated speech using Python TTS libraries.

Extensible code structure
Easily extend Tren with custom commands, functions, and tools.

Repository Structure
Tren/
│
├── app.py               # Main voice assistant logic
├── main1.py             # Additional test/experimental entry file
├── rag.py               # RAG (Retrieval-Augmented Generation) helper (if used)
├── requirements.txt     # All Python dependencies
├── LICENSE              # MIT License

Installation
1. Clone the repository
git clone https://github.com/Akash4467/Tren
cd Tren

2. Create a virtual environment (recommended)
python -m venv venv
source venv/bin/activate       # Linux & macOS
venv\Scripts\activate          # Windows

3. Install dependencies
pip install -r requirements.txt


Make sure you have:

Python 3.9 or above

Working microphone

Optional: API key if using an external LLM

Usage
Run the assistant
python app.py


or

python main1.py

How it works (flow)

Tren activates microphone listening.

Converts spoken input → text using Speech Recognition.

Sends text to an LLM (local or remote).

Receives response → converts to speech (TTS).

Plays audio output to the user.

Configuration

If your project uses an external LLM API (OpenAI, Gemini, Groq, etc.), set environment variables:

export API_KEY="your_api_key"


or create a .env file (if supported):

API_KEY=your_api_key


If RAG is enabled (rag.py), add PDFs/text files to your configured data folder.

Example Commands You Can Try

“What’s the time right now?”

“Tell me a quick summary of machine learning.”

“Explain how hashing works.”

“Open YouTube.”

“Convert 5 kilometers into meters.”

Troubleshooting
Microphone not detected

Install PyAudio properly based on OS:

pip install pyaudio


Windows users may need whl files:
https://www.lfd.uci.edu/~gohlke/pythonlibs/#pyaudio

Assistant not speaking

Check your TTS engine installation:

pip install pyttsx3

Import errors

Ensure virtual environment is active and dependencies are installed.

Future Improvements

Hotword activation (“Hey Tren…”)

Desktop automation (open apps, control keyboard/mouse)

Browser search integration

Local embeddings for offline RAG

GUI version of assistant

License

This project is licensed under the MIT License.
You are free to use, modify, and distribute the software.
