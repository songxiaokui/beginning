from django.urls import path
from chat.views.chat import chat_stream

urlpatterns = [
    path("", chat_stream, name="chat_stream"),
]