const { ZBClient } = require("zeebe-node");

(async () => {
  // const zbc = new ZBClient("localhost:26500");
  const zbc = new ZBClient();
  const topology = await zbc.topology();
  console.log(JSON.stringify(topology, null, 2));

  zbc.createWorker("collect-money", (job, complete) => {
    console.log(`[collect-money] job = ${JSON.stringify(job, null, 2)}`);
    complete.success({
      approved: true,
      customerId: 99,
    });
  });

  zbc.createWorker("fetch-items", (job, complete) => {
    console.log(`[fetch-items] job = ${JSON.stringify(job, null, 2)}`);
    complete.success();
  });

  zbc.createWorker("ship-parcel", (job, complete) => {
    console.log(`[ship-parcel] job = ${JSON.stringify(job, null, 2)}`);
    complete.success();
  });
})();
