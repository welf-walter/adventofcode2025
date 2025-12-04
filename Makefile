#DAYS = 0 1 2 3
DAYS = $(subst cmd/day,,$(wildcard cmd/day*))

#test:
#	echo $(DAYS)

all: out/day0 out/day1 out/day2 out/day3
	echo All days implemented: $(DAYS)

.PHONY : clean
clean:
	rm -rf out/

out:
	mkdir out

# use variable $(1) day number
define MAKEDAY
out/day$(1): cmd/day$(1)/day$(1).go cmd/day$(1)/* out
	echo "### Testing Day $(1) ###"
	go test adventofcode/year2025/cmd/day$(1)
	echo "### Running Day $(1) ###"
	go run $$< > $$@
	cat $$@
endef

$(foreach DAY,$(DAYS),$(eval $(call MAKEDAY,$(DAY))))
