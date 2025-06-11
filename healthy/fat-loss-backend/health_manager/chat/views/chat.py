# chat/views.py
from django.http import StreamingHttpResponse, JsonResponse
from django.views.decorators.csrf import csrf_exempt
from chat.utils.client_register import singleton, ModelType
from langchain.schema import HumanMessage
import json


@csrf_exempt
def chat_stream(request):
    if request.method != "POST":
        return JsonResponse({"error": "Only POST allowed."}, status=405)

    data = json.loads(request.body)

    def event_stream():
        data_stream = singleton.get_instance().get(
            ModelType.DeepSeek.value).stream([
            HumanMessage(content=data.get("message"))
        ])

        buffer = ""
        for chunk in data_stream:
            content = chunk.content or ""
            buffer += content

            if any(p in content for p in "，。：；！？\n") or len(buffer) > 6:
                yield buffer
                buffer = ""

        # 输出剩余部分
        if buffer:
            yield buffer

    return StreamingHttpResponse(
        event_stream(),
        content_type='text/event-stream')
