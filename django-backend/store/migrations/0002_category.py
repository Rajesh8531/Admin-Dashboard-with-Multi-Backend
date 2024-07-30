# Generated by Django 5.0.7 on 2024-07-22 12:18

import django.db.models.deletion
from django.db import migrations, models


class Migration(migrations.Migration):

    dependencies = [
        ('store', '0001_initial'),
    ]

    operations = [
        migrations.CreateModel(
            name='Category',
            fields=[
                ('id', models.CharField(max_length=50, primary_key=True, serialize=False)),
                ('createdAt', models.DateTimeField(auto_now_add=True)),
                ('updatedAt', models.DateTimeField(auto_now=True)),
                ('name', models.CharField(max_length=50)),
                ('billboardId', models.ForeignKey(db_column='billboardId', on_delete=django.db.models.deletion.CASCADE, related_name='categories', to='store.billboard')),
                ('storeId', models.ForeignKey(db_column='storeId', on_delete=django.db.models.deletion.PROTECT, related_name='categories', to='store.store')),
            ],
        ),
    ]
