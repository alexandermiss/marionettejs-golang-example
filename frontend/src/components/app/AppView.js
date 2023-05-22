import Backbone from "backbone";
import Marionette from "backbone.marionette";
import LayoutView from "src/components/app/LayoutView";
import MainRouter from "src/MainRouter";

export default Marionette.Application.extend({
  region: "#app",

  onStart() {
    const appView = new LayoutView();
    const mainAppRouter = new MainRouter({
      appView: appView,
    });
    this.showView(appView);
    Backbone.history.start();
  },
});
