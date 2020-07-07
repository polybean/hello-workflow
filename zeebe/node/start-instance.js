const { ZBClient } = require("zeebe-node");

(async () => {
  const zbc = new ZBClient();
  const result = await zbc.createWorkflowInstance("order-process", {
    customerId: 11,
  });
  console.log(result);
})();
