using System.Net;
using Nicechain;

// Helper method for key import
static byte[] StringToByteArray(string hex)
{
    return Enumerable.Range(0, hex.Length)
                     .Where(x => x % 2 == 0)
                     .Select(x => Convert.ToByte(hex.Substring(x, 2), 16))
                     .ToArray();
}

// Reads the privateKey as a result of `niceChaind keys export alice --unsafe --unarmored-hex`
// and grpcURL of the chain from the commandline
byte[] privateKey = StringToByteArray(args[0]);
String grpcURL = args[1];

// queryClient for querying, can be used standalone
var queryClient = new QueryClient(grpcURL);
// txClient for transactions
var txClient = new TxClient(queryClient, privateKey);

// Creates a new TestItem on the blockchain
await txClient.NicechainV1TxClient.SendMsgCreateTestItem(
    new Nicechain.Nicechain.V1.MsgCreateTestItem
    {
        Creator = txClient.Ec.AccoutAddress.Bech32
    }
);

// Queries all testitems
Nicechain.Nicechain.V1.QueryAllTestItemResponse response = queryClient.NicechainV1QueryClient.ListTestItem(new Nicechain.Nicechain.V1.QueryAllTestItemRequest { });
Console.Out.WriteLine(response.ToString());

// Makes sure the right amount of testItems exists on the chain now
int itemCount = response.TestItem.Count;
if (itemCount != 1)
{
    throw new Exception("expected one TestItem in chain, but got: " + itemCount);
}
