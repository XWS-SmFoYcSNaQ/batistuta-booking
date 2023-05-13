using System;
using Microsoft.EntityFrameworkCore.Migrations;

#nullable disable

#pragma warning disable CA1814 // Prefer jagged arrays over multidimensional

namespace user_service.data.Migrations
{
    /// <inheritdoc />
    public partial class InitMigration : Migration
    {
        /// <inheritdoc />
        protected override void Up(MigrationBuilder migrationBuilder)
        {
            migrationBuilder.AlterDatabase()
                .Annotation("MySQL:Charset", "utf8mb4");

            migrationBuilder.CreateTable(
                name: "Users",
                columns: table => new
                {
                    Id = table.Column<Guid>(type: "char(36)", nullable: false),
                    Role = table.Column<string>(type: "varchar(32)", maxLength: 32, nullable: false),
                    Username = table.Column<string>(type: "varchar(64)", maxLength: 64, nullable: false),
                    Password = table.Column<string>(type: "varchar(256)", maxLength: 256, nullable: false),
                    FirstName = table.Column<string>(type: "varchar(32)", maxLength: 32, nullable: false),
                    LastName = table.Column<string>(type: "varchar(32)", maxLength: 32, nullable: false),
                    Email = table.Column<string>(type: "varchar(64)", maxLength: 64, nullable: false),
                    LivingPlace = table.Column<string>(type: "varchar(64)", maxLength: 64, nullable: false)
                },
                constraints: table =>
                {
                    table.PrimaryKey("PK_Users", x => x.Id);
                    table.UniqueConstraint("AK_Users_Email", x => x.Email);
                    table.UniqueConstraint("AK_Users_Username", x => x.Username);
                })
                .Annotation("MySQL:Charset", "utf8mb4");

            migrationBuilder.InsertData(
                table: "Users",
                columns: new[] { "Id", "Email", "FirstName", "LastName", "LivingPlace", "Password", "Role", "Username" },
                values: new object[,]
                {
                    { new Guid("00695ec1-8ed0-4781-a432-9fdd98bd6cde"), "nonauthenticated@test.com", "Non", "Authenticated", "Beograd, Srbija", "AMhtGKXzeetCyumUykEQ2R0YxxLzYTJPzeQbaLQeseWBqf8bQiEmcUQV/0XRUkvQUA==", "Nonauthenticated", "nonauthenticated" },
                    { new Guid("422bd1e1-6d70-42cc-adc6-89a09b313c01"), "guest@test.com", "Guest", "Guest", "Novi Sad, Srbija", "AHD4mlEmnCAhyWDp4D6H3lcLzr5WVROCnshsLiuitVapL7GWNshdiY4pFElnRqN0qQ==", "Guest", "guest" },
                    { new Guid("87abe34c-9935-44c4-aad5-af82f9442c77"), "host@test.com", "Host", "Host", "Sabac, Srbija", "AEYedJWK/IHAD3NNv03runGcGLeN1CDKyKZS8ni/3x3gFKP8AwN5m+F2vX2kH/McjQ==", "Host", "host" }
                });
        }

        /// <inheritdoc />
        protected override void Down(MigrationBuilder migrationBuilder)
        {
            migrationBuilder.DropTable(
                name: "Users");
        }
    }
}
