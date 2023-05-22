import Marionette from "backbone.marionette";
import layout from "src/templates/layout.jst";
import NavbarView from "src/components/app/NavbarView";
import HeaderView from "src/components/app/HeaderView";

export default Marionette.View.extend({
  template: layout,
  regions: {
    nav: "nav",
    header: "header",
    main: "main",
  },
  onRender() {
    this.showChildView("nav", new NavbarView());
    this.showChildView("header", new HeaderView());
    // this.showChildView("main", new MainView());
  },
});
