const {HelloReply, HelloRequest} = require('./helloworld_pb.js');
const {GreeterClient} = require('./helloworld_grpc_web_pb.js');

const defaultEndpoint = 'http://35.200.18.55:8080/server';

const app = new Vue({
  el: '#app',
  data: {
    req: {
      endpoint: defaultEndpoint,
      message: ''
    },
    resp: {
      message: ''
    }
  },
  methods: {
    greet: function(event) {
      this.client = new GreeterClient(this.req.endpoint);

      const req = new HelloRequest();
      req.setName(this.req.message);

      this.client.sayHello(req, {}, (err, resp) => {
        if (err) {
          console.log(err.code);
          console.log(err.message);
        } else {
          console.log(resp.getMessage());
          this.resp.message = resp.getMessage();
        }
      });
    }
  }
});
