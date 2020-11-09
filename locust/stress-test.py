#!/usr/bin/env python3
# -*- coding: utf-8 -*-
''' STRESS TEST for Katalog (Goods) '''

import os
import random
import requests
from locust import User, HttpUser, TaskSet, task, constant, between

random.random()

class GoodsTest(TaskSet):
  ''' Test Products '''
  wait_time = between(1, 2)

  def on_start(self):
    """ on_start is called when a Locust start before any task is scheduled """

  @task
  def product4sale(self):
    ''' check GET to "/product4sale" '''
    self.client.get("/product4sale/%d/%d" % (random.randrange(0, 200000), random.randrange(0, 1000)))

class WebsiteUser(HttpUser):
  ''' Main class '''
  tasks = {GoodsTest:2}

  @task
  def index_page(self):
    pass
