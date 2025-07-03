import io

from fastapi import FastAPI, UploadFile, File
from fastapi.staticfiles import StaticFiles
from minio import Minio
from fastapi.responses import StreamingResponse
from uuid import uuid4
import uvicorn
from urllib.parse import quote
from config import settings

app = FastAPI(debug=settings.Debug)

app.mount("/html", StaticFiles(directory="html"), name="html")

minio_client = Minio(
    endpoint=settings.Endpoint,
    access_key=settings.AccessKey,
    secret_key=settings.SecretKey,
    secure=settings.Secure
)

BUCKET_NAME = "my-bucket"


@app.post("/upload")
async def upload(file: UploadFile = File(...)):
    object_name = f"{uuid4()}_{file.filename}"
    content = await file.read()
    # 确保 bucket 存在
    found = minio_client.bucket_exists(BUCKET_NAME)
    if not found:
        minio_client.make_bucket(BUCKET_NAME)
    # 上传
    minio_client.put_object(
        BUCKET_NAME,
        object_name,
        data=io.BytesIO(content),
        length=file.size,
        content_type=file.content_type
    )
    return {"file_id": object_name}


@app.get("/download/{file_id}")
async def download(file_id: str):
    try:
        response = minio_client.get_object(BUCKET_NAME, file_id)
        file_name_encoded = quote(file_id)
        return StreamingResponse(
            response,
            media_type="application/octet-stream",
            headers={
                "Content-Disposition": f"attachment; filename={file_name_encoded}"
            }
        )
    except Exception as e:
        return {"error": str(e)}


if __name__ == "__main__":
    uvicorn.run(app, host="0.0.0.0", port=9115)
