from chromadb.api.types import Documents, EmbeddingFunction, Embeddings


class SentenceTransformerEmbeddingFunction(EmbeddingFunction):
    def __init__(self, model_name: str = "all-MiniLM-L6-v2"):
        try:
            from sentence_transformers import SentenceTransformer
        except ImportError:
            raise ValueError(
                "The sentence_transformers python package is not installed. Please install it with `pip install sentence_transformers`"
            )
        self._model = SentenceTransformer(model_name)

    def __call__(self, texts: Documents) -> Embeddings:
        return self._model.encode(list(texts), convert_to_numpy=True).tolist()


class OpenAIEmbeddingFunction(EmbeddingFunction):
    def __init__(self, api_key: str, model_name: str = "text-embedding-ada-002"):
        try:
            import openai
        except ImportError:
            raise ValueError(
                "The openai python package is not installed. Please install it with `pip install openai`"
            )

        openai.api_key = api_key
        self._client = openai.Embedding
        self._model_name = model_name

    def __call__(self, texts: Documents) -> Embeddings:
        # replace newlines, which can negatively affect performance.
        texts = [t.replace("\n", " ") for t in texts]
        # Call the OpenAI Embedding API in parallel for each document
        return [
            result["embedding"]
            for result in self._client.create(
                input=texts,
                engine=self._model_name,
            )["data"]
        ]


class CohereEmbeddingFunction(EmbeddingFunction):
    def __init__(self, api_key: str, model_name: str = "large"):
        try:
            import cohere
        except ImportError:
            raise ValueError(
                "The cohere python package is not installed. Please install it with `pip install cohere`"
            )

        self._client = cohere.Client(api_key)
        self._model_name = model_name

    def __call__(self, texts: Documents) -> Embeddings:
        # Call Cohere Embedding API for each document.
        return [
            embeddings for embeddings in self._client.embed(texts=texts, model=self._model_name)
        ]

class HuggingFaceEmbeddingFunction(EmbeddingFunction):
    def __init__(self, api_key: str, model_name: str = "sentence-transformers/all-MiniLM-L6-v2"):
        try:
            import requests
        except ImportError:
            raise ValueError(
                "The requests python package is not installed. Please install it with `pip install requests`"
            )
        self.api_url = f"https://api-inference.huggingface.co/pipeline/feature-extraction/{model_id}"
        self._headers = {"Authorization": f"Bearer {hf_token}"}
        self._model_name = model_name

    def __call__(self, texts: Documents) -> Embeddings:
        # Call HuggingFace Embedding API for each document.
        return requests.post(api_url, headers=headers, json={"inputs": texts, "options":{"wait_for_model":True}}).json()
