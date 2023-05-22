import Marionette from "backbone.marionette";
import provider_table from "src/templates/providers/provider_table.jst";
import ProviderListItem from "src/components/providers/ProviderListItem";

export default Marionette.CollectionView.extend({
  template: provider_table,
  region: "main",
  childViewContainer: "tbody",
  childView: ProviderListItem,
});
