import Marionette from "backbone.marionette";
import client_add_form from "src/templates/clients/client_add_form.jst";

export default Marionette.View.extend({
  template: client_add_form,
  region: "main",
  // initialize(opts) {
  //   this.collection = opts.collection;
  // },
  serializeData() {
    const data = _.clone(this.collection.toJSON());
    return {
      providers: data,
    };
  },
});
