from django.db import models
from .base import TimeStampedModel


class CheckIn(TimeStampedModel):
    _id = models.AutoField(primary_key=True)
    user_id = models.CharField(max_length=64)
    date = models.DateField(unique=False)
    breakfast = models.BooleanField(default=False)
    lunch = models.BooleanField(default=False)
    dinner = models.BooleanField(default=False)
    exercise = models.BooleanField(default=False)
    sleep = models.BooleanField(default=False)

    class Meta:
        db_table = "checkin"

    def __str__(self):
        return f"{self.user_id} | {self.date} 打卡"
