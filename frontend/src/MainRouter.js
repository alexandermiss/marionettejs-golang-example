import Backbone from "backbone";
import ClientListView from "src/components/clients/ClientListView";
import ClientDetailView from "src/components/clients/ClientDetailView";

import ProviderListView from "src/components/providers/ProviderListView";
import ProviderDetailView from "src/components/providers/ProviderDetailView";

import AddFormClientView from "src/components/clients/AddFormClientView";

import LayoutView from "src/components/app/LayoutView";

export default Backbone.Router.extend({
  initialize(opts) {
    this.appView = opts.appView;
  },
  routes: {
    "": "clientList",
    clients: "clientList",
    "clients/add": "addClientView",
    "clients/:id": "showClientDetail",
    providers: "showProviderListView",
    "providers/add": "addProviderView",
    "providers/:id": "showProviderDetail",
  },

  clientList() {
    const UserModel = Backbone.Model.extend({
      idAttribute: "ID",
    });

    const UserCollection = Backbone.Collection.extend({
      url: "http://localhost:8599/api/v1/client",
    });

    const clients = new UserCollection();
    clients.fetch({ reset: true });

    const clientListView = new ClientListView({ collection: clients });
    this.appView.getRegion("main").show(clientListView);

    const headerView = this.appView.getChildView("header");
    headerView.getUI("title").html("Clients");
  },
  showClientDetail(id) {
    const UserModel = Backbone.Model.extend({
      idAttribute: "ID",
      urlRoot: "http://localhost:8599/api/v1/client",
    });

    const model = new UserModel({ ID: id });
    model.fetch({ reset: true });

    var self = this;
    model.on("sync", function () {
      const clientDetailView = new ClientDetailView({
        model: model,
      });

      self.appView.getRegion("main").show(clientDetailView);

      const headerView = self.appView.getChildView("header");
      headerView.getUI("title").html("Client #" + model.get("ID"));
    });
  },
  showProviderListView() {
    const Model = Backbone.Model.extend({
      idAttribute: "ID",
    });

    const ProviderCollection = Backbone.Collection.extend({
      url: "http://localhost:8599/api/v1/provider",
      model: Model,
    });

    const collection = new ProviderCollection();
    collection.fetch({ reset: true });

    const providerListView = new ProviderListView({ collection: collection });
    this.appView.getRegion("main").show(providerListView);

    const headerView = this.appView.getChildView("header");
    headerView.getUI("title").html("Providers");
  },
  showProviderDetail(id) {
    const ProviderModel = Backbone.Model.extend({
      idAttribute: "ID",
      urlRoot: "http://localhost:8599/api/v1/provider",
    });

    const model = new ProviderModel({ ID: id });
    model.fetch({ reset: true });

    var self = this;
    model.on("sync", function () {
      const providerDetailView = new ProviderDetailView({
        model: model,
      });

      self.appView.getRegion("main").show(providerDetailView);

      const headerView = self.appView.getChildView("header");
      headerView.getUI("title").html("Provider #" + model.get("ID"));
    });
  },
  addClientView() {
    const Model = Backbone.Model.extend({
      idAttribute: "ID",
    });

    const ProviderCollection = Backbone.Collection.extend({
      url: "http://localhost:8599/api/v1/provider",
      model: Model,
    });

    const collection = new ProviderCollection();
    collection.fetch({ reset: true });

    var self = this;
    collection.on("sync", function () {
      const addFormClientView = new AddFormClientView({
        collection: collection,
      });
      self.appView.getRegion("main").show(addFormClientView);
    });
  },
  addProviderView() {
    this.appView.getRegion("main").show(providerDetailView);
  },
});
