const { ZBClient } = require("zeebe-node");

(async () => {
  const zbc = new ZBClient({
    onConnectionError: (err) => console.log("err", err),
    onReady: () => console.log("YOO"),
  });
  const topology = await zbc.topology();
  console.log(JSON.stringify(topology, null, 2));

  const res = await zbc.deployWorkflow("./order-process.bpmn");
  setTimeout(() => console.log(res), 5000);
})();
