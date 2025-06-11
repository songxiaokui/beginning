# chat/views.py
from django.http import StreamingHttpResponse, JsonResponse
import time
from django.views.decorators.csrf import csrf_exempt


@csrf_exempt
def chat_stream(request):
    if request.method != "POST":
        return JsonResponse({"error": "Only POST allowed."}, status=405)

    def event_stream():
        # 模拟流式输出
        yield "Hello"
        time.sleep(0.5)
        yield "\n\n, "
        time.sleep(0.5)
        yield "\n\nWorld!"
        time.sleep(0.5)
        yield "\n\n"

    return StreamingHttpResponse(event_stream(), content_type='text/event-stream')
