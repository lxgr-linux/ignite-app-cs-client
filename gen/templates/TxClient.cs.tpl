using System.Net.Http;
using Cosmcs.Client;

namespace {{ .NameSpace }}
{
    public class TxClient
    {
        public EasyClient Ec { get; }
        public QueryClient QueryClient { get; }
        {{- range .Services }}
        public {{ .Path }}.{{ .Type }}Client {{ .Name }}TxClient { get; }
        {{- end }}

        public TxClient(QueryClient queryClient, byte[] privateKey)
        {
            Ec = new EasyClient(queryClient, privateKey);
            QueryClient = queryClient;
            {{- range .Services }}
            {{ .Name }}TxClient = new {{ .Path }}.{{ .Type }}Client(Ec);
            {{- end }}
        }
    }
}
