from django.contrib import admin
from checkin.models import CheckIn


@admin.register(CheckIn)
class CheckInAdmin(admin.ModelAdmin):
    list_display = (
        "user_id",
        "date",
        "breakfast",
        "lunch",
        "dinner",
        "exercise",
        "sleep",
        "created_at",
    )
    list_filter = ("date", "user_id")
    search_fields = ("user_id",)
