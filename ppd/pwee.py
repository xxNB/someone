from peewee import MySQLDatabase, Model, IntegerField, FloatField, CharField

db = MySQLDatabase(
    "data",  # Required by Peewee.
    user="root",  # Will be passed directly to psycopg2.
    # password='secret',  # Ditto.
    host="127.0.0.1",
    charset="utf8mb4",
)  # Ditto.

class Alldue(Model):
    installmentype = IntegerField(index=True)
    scorebin = IntegerField(index=True)
    rulebin = IntegerField(index=True)
    overduevalue = FloatField()


    class Meta:
        database = db # This model uses the "people.db" database.

class EveryDue(Model):
    installmentype = IntegerField(index=True)
    scorebin = IntegerField(index=True)
    rulebin = IntegerField(index=True)
    loanperiod = IntegerField()
    radio = FloatField()

    class Meta:
        database = db # This model uses the "people.db" database.

if __name__ == "__main__":
    db.connect()
    db.create_tables([Alldue,EveryDue])