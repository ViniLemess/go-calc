package com.vinilemess.simulation.gocalc

import io.gatling.core.Predef._
import io.gatling.http.Predef._
import scala.concurrent.duration._

class GoCalcGetHistorySimulation extends Simulation {
  val httpProtocol = http.baseUrl("http://localhost:8090")

  val historyScn = scenario("History Stress Test")
    .exec(http("Http Round Robin 1")
      .get("/calc/history")
      .check(status.is(200))
    )

  setUp(
    historyScn.inject(
      nothingFor(4),
      atOnceUsers(10),
      rampUsers(10).during(5),
      constantUsersPerSec(20).during(15),
      constantUsersPerSec(20).during(15).randomized,
      rampUsersPerSec(10).to(20).during(10.seconds),
      rampUsersPerSec(10).to(20).during(10.seconds).randomized,
      constantUsersPerSec(1000).during(3)
    ).protocols(httpProtocol)
  ).assertions(global.responseTime.max.lt(1000), global.successfulRequests.percent.gte(95))
}
