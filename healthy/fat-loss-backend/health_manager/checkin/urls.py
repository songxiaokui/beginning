from django.urls import path
from checkin.views.checkin_views import get_checkin, post_checkin

urlpatterns = [
    path("checkin", get_checkin, name="get_checkin"),  # GET /api/checkin
    path(
        "checkin/submit", post_checkin, name="post_checkin"
    ),  # POST /api/checkin/submit
]
