using System;
using Microsoft.EntityFrameworkCore.Migrations;

#nullable disable

#pragma warning disable CA1814 // Prefer jagged arrays over multidimensional

namespace user_service.data.Migrations
{
    /// <inheritdoc />
    public partial class InitMigrationWithSeedData : Migration
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
                    { new Guid("422bd1e1-6d70-42cc-adc6-89a09b313c01"), "guest@test.com", "Guest", "Guest", "Novi Sad, Srbija", "10000.GhMJYLVMJDUSKPYAt3G+oA==.2d2SyAT1CWcY/eNqJiKXKdTvjrWY2TftfJsHiOCy54g=", "Guest", "guest" },
                    { new Guid("87abe34c-9935-44c4-aad5-af82f9442c77"), "host@test.com", "Host", "Host", "Sabac, Srbija", "10000.pUrW8b1z1nt7+RFCVYWpWg==.jp+r7PJ49rgJwwZVAfLhCb2YMyCAZR3gXgrLnson2UQ=", "Host", "host" }
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
