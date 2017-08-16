const log = require('../utils/log');
const installerInput = require('../utils/bareMetalInstallerInput');
const tfvarsUtil = require('../utils/terraformTfvars');

module.exports = {
  after (client) {
    client.getLog('browser', log.logger);
    client.end();
  },

  'Tectonic Installer BareMetal Test': (client) => {
    const expectedJson = installerInput.buildExpectedJson();
    const platformPage = client.page.platformPage();
    const clusterInfoPage = client.page.clusterInfoPage();
    const clusterDnsPage = client.page.clusterDnsPage();
    const certificateAuthorityPage = client.page.certificateAuthorityPage();
    const matchboxAddressPage = client.page.matchboxAddressPage();
    const matchboxCredentialsPage = client.page.matchboxCredentialsPage();
    const defineMastersPage = client.page.defineMastersPage();
    const defineWorkersPage = client.page.defineWorkersPage();
    const etcdConnectionPage = client.page.etcdConnectionPage();
    const networkConfigurationPage = client.page.networkConfigurationPage();
    const sshKeysPage = client.page.sshKeysPage();
    const consoleLoginPage = client.page.consoleLoginPage();
    const submitPage = client.page.submitPage();

    platformPage.navigate(client.launch_url).selectBareMetalPlatform();
    clusterInfoPage.enterClusterInfo(expectedJson.tectonic_cluster_name = `baremetaltest-${new Date().getTime().toString()}`);
    clusterDnsPage.enterDnsNames();
    certificateAuthorityPage.click('@nextStep');
    matchboxAddressPage.enterMatchBoxEndPoints();
    matchboxCredentialsPage.enterMatchBoxCredentials();
    networkConfigurationPage.enterCIDRs();
    defineMastersPage.enterMastersDnsNames();
    defineWorkersPage.enterWorkersDnsNames();
    etcdConnectionPage.click('@nextStep');
    sshKeysPage.enterPublicKey();
    consoleLoginPage.enterLoginCredentails(expectedJson.tectonic_admin_email);
    submitPage.click('@manuallyBoot');
    client.pause(10000);
    client.getCookie('tectonic-installer', result => {
      tfvarsUtil.returnTerraformTfvars(client.launch_url, result.value, (err, actualJson) => {
        if (err) {
          return client.assert.fail(err);
        }
        const msg = tfvarsUtil.compareJson(actualJson, expectedJson);
        if (msg) {
          return client.assert.fail(msg);
        }
      });
    });
  },
};
