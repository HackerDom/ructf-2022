using System;
using System.Collections.Generic;
using System.Diagnostics.CodeAnalysis;
using System.IO;
using System.Linq;
using System.Text;
using System.Threading;
using Newtonsoft.Json;
using Newtonsoft.Json.Converters;
using Newtonsoft.Json.Linq;

namespace psycho_clinic.Extensions
{
    internal static class JsonSerialization
    {
        public static readonly JsonSerializerSettings Settings;
        public static readonly JsonSerializer PrimitiveSerializer;
        private static readonly IReadOnlyList<JsonConverter> Converters;
        private static readonly JsonSerializerSettings ErrorToleranceSettings;

        private static readonly JsonSerializer Serializer;

        private static readonly ThreadLocal<bool> HasDeserializationError;

        static JsonSerialization()
        {
            HasDeserializationError = new(() => false);

            IReadOnlyList<JsonConverter> primitiveConverters = new List<JsonConverter>
            {
                // new IPAddressJsonConverter(),
                new StringEnumConverter(),
                new VersionConverter(),
                // new ComplexDictionaryJsonConverter(),
                // new ComparableDictionaryJsonConverter(),
                // new ComparableArrayJsonConverter(),
                // new DeploymentLogEventScopeJsonConverter(),
                // new StringSerializationJsonConverter(),
                // new ReadOnlySetJsonConverter(),
                // new DeploymentSettingsBackwardCompatibilityJsonConverter(),
                // new GroupDeploymentSettingsBackwardCompatibilityJsonConverter(),
            };
            Converters = primitiveConverters;

            Settings = SetupCommonSettings(new());
            Serializer = JsonSerializer.CreateDefault(Settings);

            var primitiveSettings = SetupCommonSettings(new());
            primitiveSettings.Converters = primitiveConverters.ToList();
            PrimitiveSerializer = JsonSerializer.CreateDefault(primitiveSettings);

            ErrorToleranceSettings = SetupCommonSettings(new()
            {
                Error = (_, args) =>
                {
                    HasDeserializationError.Value = true;
                    args.ErrorContext.Handled = true;
                },
                MissingMemberHandling = MissingMemberHandling.Error
            });
        }

        public static string ToJson(this object? @object) => JsonConvert.SerializeObject(@object, Settings);

        public static JObject ToJObject(this object @object) => JObject.FromObject(@object, Serializer);

        public static T FromJson<T>(this string serialized) => JsonConvert.DeserializeObject<T>(serialized, Settings)!;

        public static T FromJsonStream<T>(this Stream serialized)
        {
            using var streamReader = new StreamReader(serialized);
            using var jsonReader = new JsonTextReader(streamReader);

            return Serializer.Deserialize<T>(jsonReader)!;
        }

        public static object? FromJson(this string serialized, Type type) =>
            JsonConvert.DeserializeObject(serialized, type, Settings);

        public static object? FromJson(this string serialized) => JsonConvert.DeserializeObject(serialized, Settings);

        [return: MaybeNull]
        public static T FromJObject<T>(this JObject serialized) => serialized.ToObject<T>(Serializer);

        public static T? FromJson<T>(this byte[]? content)
            where T : class
        {
            return content == null
                ? null
                : Encoding.UTF8.GetString(content).FromJson<T>();
        }

        public static bool TryFromJson<T>(this string serialized, out T? deserialized)
            where T : class
        {
            HasDeserializationError.Value = false;
            deserialized = JsonConvert.DeserializeObject<T>(serialized, ErrorToleranceSettings);
            return !HasDeserializationError.Value;
        }

        public static T CloneJson<T>(this T @object) => @object.ToJson().FromJson<T>();

        public static JsonSerializerSettings SetupCommonSettings(JsonSerializerSettings settings)
        {
            settings.NullValueHandling = NullValueHandling.Ignore;
            settings.DateParseHandling = DateParseHandling.None;

            foreach (var converter in Converters)
                settings.Converters.Add(converter);

            return settings;
        }
    }
}