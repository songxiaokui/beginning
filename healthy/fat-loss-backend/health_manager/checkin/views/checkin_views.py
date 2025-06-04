from django.http import JsonResponse
from django.views.decorators.csrf import csrf_exempt
from checkin.models import CheckIn
import json
from datetime import date


@csrf_exempt
def get_checkin(request):
    q_date = request.GET.get("date", date.today().isoformat())
    user_id = request.GET.get("user_id")
    if not user_id:
        return JsonResponse({"error": "missing user_id"}, status=400)

    try:
        record = CheckIn.objects.get(user_id=user_id, date=q_date)
        return JsonResponse(
            {
                "breakfast": record.breakfast,
                "lunch": record.lunch,
                "dinner": record.dinner,
                "exercise": record.exercise,
                "sleep": record.sleep,
            }
        )
    except CheckIn.DoesNotExist:
        return JsonResponse(
            {
                "breakfast": False,
                "lunch": False,
                "dinner": False,
                "exercise": False,
                "sleep": False,
            }
        )


@csrf_exempt
def post_checkin(request):
    if request.method != "POST":
        return JsonResponse({"error": "Only POST method allowed"}, status=405)

    try:
        data = json.loads(request.body)
    except json.JSONDecodeError:
        return JsonResponse({"error": "Invalid JSON"}, status=400)

    q_date = data.get("data", date.today().isoformat())
    user_id = data.get("user_id")
    config = data.get("config", {})

    if not user_id:
        return JsonResponse({"error": "missing user_id"}, status=400)

    record, _ = CheckIn.objects.get_or_create(user_id=user_id, date=q_date)
    record.breakfast = config.get("breakfast", False)
    record.lunch = config.get("lunch", False)
    record.dinner = config.get("dinner", False)
    record.exercise = config.get("exercise", False)
    record.sleep = config.get("sleep", False)
    record.save()

    return JsonResponse({"message": "Check-in saved"})
