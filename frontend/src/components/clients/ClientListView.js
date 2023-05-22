import Marionette from "backbone.marionette";
import client_table from "src/templates/clients/client_table.jst";
import ClientListItem from "src/components/clients/ClientListItemView";

export default Marionette.CollectionView.extend({
  template: client_table,
  region: "main",
  childViewContainer: "tbody",
  childView: ClientListItem,
});
